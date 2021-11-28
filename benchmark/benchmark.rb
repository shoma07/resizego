# frozen_string_literal: true

require "bundler/setup"
require "benchmark"
require "posix-spawn"
require "mini_magick"
require "resizego"

MiniMagick.configure do |config|
  config.shell_api = "posix-spawn"
end

Benchmark.bm do |x|
  x.report("mini_magick") do
    Dir["#{__dir__}/images/*"].each do |path|
      image = MiniMagick::Image.create("jpg", false) do |tempfile|
        IO.copy_stream(StringIO.new(File.binread(path)), tempfile)
      end
      image.combine_options do |b|
        b.auto_orient
        b.strip
        b.resize("600x600")
      end.to_blob
    end
  end
  x.report("resizego") do
    Dir["#{__dir__}/images/*"].each do |path|
      Resizego.resize(File.binread(path), 600)
    end
  end
end
