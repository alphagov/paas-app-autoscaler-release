require 'rspec'
require 'json'
require 'bosh/template/test'
require 'yaml'

describe 'metricsforwarder' do
  let(:release) { Bosh::Template::Test::ReleaseDir.new(File.join(File.dirname(__FILE__), '../../..')) }
  let(:job) { release.job('metricsforwarder') }
  let(:template) { job.template('config/metricsforwarder.yml') }
  let(:properties) do
    YAML.safe_load(%(
      autoscaler:
        policy_db:
          address: 10.11.137.101
          databases:
          - name: foo
            tag: default
          db_scheme: postgres
          port: 5432
          roles:
          - name: foo
            password: default
            tag: default
        cf:
          api: https://api.cf.domain
          auth_endpoint: https://login.cf.domain
          client_id: client_id
          secret: uaa_secret
          uaa_api: https://login.cf.domain/uaa
          grant_type: ALLOW_ALL
    ))
  end

  context 'config/metricsforwarder.yml' do

    it 'does not set username nor password if not configured' do
      properties['autoscaler'].merge!(
        'metricsforwarder' => {
          'health' => {
            'port' => 1234,
          }
        }
      )
      rendered_template = YAML.safe_load(template.render(properties))

      expect(rendered_template['health']).
        to include(
             { 'port' => 1234 }
           )
    end

    it 'check metricsforwarder basic auth username and password' do
      properties['autoscaler'].merge!(
        'metricsforwarder' => {
          'health' => {
            'port' => 1234,
            'username' => 'test-user',
            'password' => 'test-user-password'
          }
        }
      )
      rendered_template = YAML.safe_load(template.render(properties))

      expect(rendered_template['health']).
        to include(
             { 'port' => 1234,
               'username' => 'test-user',
               'password' => 'test-user-password'
             }
           )
		end

    it 'has a cred helper plugin by default' do
      rendered_template = YAML.safe_load(template.render(properties))
      expect(rendered_template).to include(
        {
          "cred_helper_plugin" => "default"
        }
      )
    end

    it 'has a cred helper plugin by configured for stored procedures' do

      properties['autoscaler'].merge!(
        'metricsforwarder' => {
          'cred_helper' => {
            'plugin' => 'stored_procedure',
            'stored_procedure_config' => {
              'schema_name' => 'SCHEMA',
              'create_binding_credential_procedure_name' => 'CREATE_BINDING_CREDENTIAL',
              'drop_binding_credential_procedure_name' => 'DROP_BINDING_CREDENTIAL',
              'drop_all_binding_credential_procedure_name' => 'DROP_ALL_BINDING_CREDENTIALS',
              'validate_binding_credential_procedure_name' => 'VALIDATE_BINDING_CREDENTIALS'
            }
          }
        }
      )

      rendered_template = YAML.safe_load(template.render(properties))
      expect(rendered_template).to include(
                                     {
                                       'cred_helper_plugin' => 'stored_procedure',
                                       'stored_procedure_config' => {
                                         'schema_name' => 'SCHEMA',
                                         'create_binding_credential_procedure_name' => 'CREATE_BINDING_CREDENTIAL',
                                         'drop_binding_credential_procedure_name' => 'DROP_BINDING_CREDENTIAL',
                                         'drop_all_binding_credential_procedure_name' => 'DROP_ALL_BINDING_CREDENTIALS',
                                         'validate_binding_credential_procedure_name' => 'VALIDATE_BINDING_CREDENTIALS'
                                       }
                                     }
                                   )
    end

    it 'has a cred helper plugin that can be configured by specifying different path' do
      properties['autoscaler'].merge!(
        'metricsforwarder' => {
          'cred_helper' => {
            'plugin' => "/var/vcap/packages/other-package-plugin"
          }
        }
      )

      rendered_template = YAML.safe_load(template.render(properties))
      expect(rendered_template).to include(
             {
               "cred_helper_plugin" => "/var/vcap/packages/other-package-plugin"
             }
      )
    end
  end
end

