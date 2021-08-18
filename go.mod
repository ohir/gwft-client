module example.com/mC

go 1.16

require example.com/mP v0.1.5 // protocol

replace example.com/mP => github.com/ohir/gwft-protocol v0.1.5
// replace example.com/mP => ../protocol
