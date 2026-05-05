#nix-shell
{pkgs ? import <nixpkgs> {}}:

pkgs.mkShell {
    nativeBuildInputs = with pkgs; [
    pkg-config
    ];

    buildInputs = with pkgs; [
      go # pin the toolchain to nixpkgs
      libGL # OpenGL loader
      alsa-lib # audio
      libX11 # X11 client x
      libXcursor # cursor extension
      libXrandr # display config
      libXinerama # multi-monitor info
      libXi # input extension
      libXxf86vm # video mode extension 
    ];

    LD_LIBRARY_PATH = pkgs.lib.makeLibraryPath (with pkgs; [
      libGL # runtime dlopen target
      libX11 # belt and suspenders
      libXcursor
      libXrandr
      libXinerama
      libXi
      libXxf86vm
    ]);
}
