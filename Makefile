.PHONY: welcome check fmt

welcome:
	@nix run .#ubctl welcome

check:
	@find ./nix -type f -name "*.sh" -exec shellcheck -o all {} \;
	@shellcheck -o all ./install.sh

fmt:
	@nix fmt

