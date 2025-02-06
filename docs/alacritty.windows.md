# Alacritty Windows Setup

1. Use the `.msi` installer on the [Alacritty releases](https://github.com/alacritty/alacritty/releases/tag/v0.13.2) page to install alacritty
1. Download a [Nerd Font](https://www.nerdfonts.com/) of your choice and make sure it is added to your font library
1. Create an Alacritty config file at `Users/<your user>/AppData/Roaming/alacritty/alacritty.toml` with the following values (the font family should be modified accordingly if you chose a font other than CaskaydiaCove):

  ```toml
  [font]
  size = 12.0 # modify as needed

  [font.bold]
  family = "CaskaydiaCove Nerd Font"
  style = "Regular"

  [font.bold_italic]
  family = "CaskaydiaCove Nerd Font"
  style = "Regular"

  [font.italic]
  family = "CaskaydiaCove Nerd Font"
  style = "Regular"

  [font.normal]
  family = "CaskaydiaCove Nerd Font"
  style = "Regular"

  [[keyboard.bindings]]
  action = "Paste"
  key = "V"
  mods = "Control"

  [[keyboard.bindings]]
  action = "Copy"
  key = "C"
  mods = "Control"

  [[keyboard.bindings]]
  chars = "\u0016"
  key = "V"
  mods = "Control|Shift"

  [[keyboard.bindings]]
  chars = "\u0003"
  key = "C"
  mods = "Control|Shift"

  [shell]
  args = ["--cd ~"]
  program = 'C:\Windows\System32\wsl.exe'
  ```

