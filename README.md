# gotk4

[![built with nix](https://builtwithnix.org/badge.svg)](https://builtwithnix.org)

A GTK4 bindings generator for Go.

Progress tracker: https://github.com/diamondburned/gotk4/issues/2

All generated packages are in `pkg/`. The generation code is in `gir/girgen/`.
At the moment, the repository depends on gotk3's GLib. This may change in the
future.

## Contributing to gotk4

For contributing guidelines, see [CONTRIBUTING.md](./CONTRIBUTING.md).

## Wishes

- I wish I generated models first before generating functions
- I wish I used a registry of converted model types with their respective Go
  names to ease translation
- I wish I used common struct types that share names and such.

## License

`gotk4` contains 3 directories licensed differently:

- `gotk4/gir` is licensed under the [GNU Affero General Public License v3][AGPLv3].
  This license does not apply to the code generated by itself.
- `gotk4/pkg` is licensed under the [Mozilla Public License v2][MPLv2].
- `gotk4/core` is licensed under the permissive [ISC][ISC] license.

[AGPLv3]: https://www.gnu.org/licenses/agpl-3.0.en.html
[MPLv2]: https://www.mozilla.org/en-US/MPL/
[ISC]: https://choosealicense.com/licenses/isc/
