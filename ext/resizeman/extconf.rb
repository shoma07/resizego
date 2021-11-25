# frozen_string_literal: true

require "fileutils"

FileUtils.cp_r("#{__dir__}/.", Dir.pwd.to_s)
