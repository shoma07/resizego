# frozen_string_literal: true

require "bundler/gem_tasks"
require "rspec/core/rake_task"

RSpec::Core::RakeTask.new(:spec)

require "rubocop/rake_task"

RuboCop::RakeTask.new

task :compile do
  system("cd #{File.expand_path(__dir__)}; ruby ext/resizego/extconf.rb")
end

task build: :compile

task default: %i[clobber compile spec rubocop]
