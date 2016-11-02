# Distributed Systems Lab One 

Niall Paterson #12305503

## Installation

 * Run `go get github.com/paterson/firstlab`
  
## Run

  * Run `firstlab` (It'll be automatically in the PATH so will work from anywhere)
  * You can change the port by supplying an argument: e.g `firstlab 3001`. It will default to `8000` otherwise.

## Details

  * The client is currently pointing to one of my VM's, which is running the following command to run the server and keep it up:

        nohup php -S 0.0.0.0:8000 -t . > php.log &
         
  
  
