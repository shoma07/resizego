# frozen_string_literal: true

require "ffi"
require "base64"
require_relative "resizego/version"

# Resizego
module Resizego
  # Resizego::Error
  class Error < StandardError; end

  extend FFI::Library

  ffi_lib File.expand_path("resizego/resizego.bundle", __dir__)

  attach_function :resize64, %i[string int int], :string

  class << self
    private :resize64

    # @param blob [String]
    # @param limit [Integer]
    # @param quality [Integer]
    # @return [String]
    # @raise [Error]
    def resize(blob, limit, quality: 100)
      Base64.decode64(resize64(Base64.encode64(blob), limit, quality))
    rescue StandardError
      raise Error, "resize error"
    end
  end
end
