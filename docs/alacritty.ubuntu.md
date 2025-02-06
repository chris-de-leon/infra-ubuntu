# Alacritty Ubuntu Setup

1. Use the App Center to install Alacritty
2. Download a [Nerd Font](https://www.nerdfonts.com/) of your choice (CaskaydiaCove is recommended)
3. Unzip the font
4. Copy the font files to `~/.fonts`
5. Run `fc-cache -fv` to update Ubuntu's font cache
6. Create an Alacritty config file at `~/.config/alacritty/alacritty.toml` with the following values (the font family should be modified accordingly if you chose a font other than CaskaydiaCove):

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
  ```

