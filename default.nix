with import <nixpkgs> {};

stdenv.mkDerivation {
  name = "mc-4.8.12";

  src = ./.;

  buildInputs = [
    xlibs.libX11 xorg.libXcursor xorg.libXrandr xorg.libXinerama mesa mesa_glu
    xorg.libXi
    # glfw xlibs.libICE
  ];

  meta = {
    description = "Graphical stl file watcher.";
  };
}
