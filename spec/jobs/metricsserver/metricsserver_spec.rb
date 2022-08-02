require "rspec"
require "json"
require "bosh/template/test"
require "yaml"
require "rspec/file_fixtures"

describe "metricsserver" do
  let(:release) { Bosh::Template::Test::ReleaseDir.new(File.join(File.dirname(__FILE__), "../../..")) }
  let(:job) { release.job("metricsserver") }
  let(:template) { job.template("config/metricsserver.yml") }
  let(:properties) { YAML.safe_load(fixture("metricsserver.yml").read) }

  context "config/metricsserver.yml" do
    it "does not set username nor password if not configured" do
      properties["autoscaler"]["metricsserver"] = {
        "health" => {
          "port" => 1234
        }
      }
      links = [
        Bosh::Template::Test::Link.new(
          name: "metricsserver"
        )
      ]
      rendered_template = YAML.safe_load(template.render(properties, consumes: links))

      expect(rendered_template["health"])
        .to include(
          {"port" => 1234}
        )
    end

    it "check metricsserver basic auth username and password" do
      properties["autoscaler"]["metricsserver"] = {
        "health" => {
          "port" => 1234,
          "username" => "test-user",
          "password" => "test-user-password"
        }
      }
      links = [
        Bosh::Template::Test::Link.new(
          name: "metricsserver"
        )
      ]
      rendered_template = YAML.safe_load(template.render(properties, consumes: links))

      expect(rendered_template["health"])
        .to include(
          {"port" => 1234,
           "username" => "test-user",
           "password" => "test-user-password"}
        )
    end
  end
end
