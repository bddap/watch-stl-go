with import <nixpkgs> {};

stdenv.mkDerivation {
  name = "watch-stl";

  src = ./.;

  buildInputs = [
    xlibs.libX11 xorg.libXcursor xorg.libXrandr xorg.libXinerama mesa mesa_glu
    xorg.libXi
  ];

  meta = {
    description = "Graphical stl file watcher.";
  };
}
