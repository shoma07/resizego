# frozen_string_literal: true

require "optparse"

module Resizego
  # Resizego::CLI
  class CLI
    # @param argv [Array]
    # @return [void]
    def initialize(argv)
      @options = { resize: 100, quality: 100 }
      parsed_argv = parser.parse!(argv)
      @blob = File.binread(parsed_argv.shift)
    end

    # @return [void]
    def run
      puts Resizego.resize(@blob, @options[:resize], quality: @options[:quality])
    end

    private

    # @return [OptionParser]
    def parser
      OptionParser.new do |opt|
        opt.on("-r", "--resize VAL", Integer) do |i|
          @options[:resize] = i
        end
        opt.on("-q", "--quality [VAL]", Integer) do |i|
          @options[:quality] = i
        end
      end
    end
  end
end
