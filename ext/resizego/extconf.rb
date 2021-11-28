# frozen_string_literal: true

require "fileutils"

pwd = Dir.pwd.to_s
path = File.expand_path(__dir__)
FileUtils.cp_r("#{path}/.", pwd) if pwd != path
