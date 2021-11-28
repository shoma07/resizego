# frozen_string_literal: true

require "fileutils"

ext_path = File.expand_path(__dir__)
root_path = File.expand_path("../../", ext_path)
tmp_path = File.expand_path("tmp/#{RUBY_PLATFORM}/resizego/#{RUBY_VERSION}", root_path)
FileUtils.mkdir_p(tmp_path)
FileUtils.cp_r("#{ext_path}/.", tmp_path)
FileUtils.cd(tmp_path)
system("make")
FileUtils.cp("#{tmp_path}/resizego.so", "#{root_path}/lib/resizego/")
