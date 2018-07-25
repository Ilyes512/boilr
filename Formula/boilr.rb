class Boilr < Formula
    desc "Boilerplate template manager that generates files or directories from template repositories"
    homepage "https://github.com/Ilyes512/boilr"
    url "https://github.com/Ilyes512/boilr/releases/download/0.4.2/boilr-0.4.2-darwin_amd64.tgz"
    version "0.4.2"
    sha256 "6dbbc1467caaab977bf4c9f2d106ceadfedd954b6a4848c54c925aff81159a65"
    head "https://github.com/Ilyes512/boilr.git"

    bottle :unneeded

    def install
        bin.install "boilr"
    end

    test do
        assert_match version.to_s, shell_output("#{bin}/boilr version --dont-prettify")
    end
end