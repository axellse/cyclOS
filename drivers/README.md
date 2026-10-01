# cyclOS drivers
This directory contains cyclOS drivers. It functions like this:
* This directory contains a number of go packages (subdirectories), one for each driver class. These contain source code files for all the different drivers, one for each class.
* Each driver implements a driver class defined by the HAL. See the related `ihal` package for these interface types as well as documentation about their expected behaviour.
* Each driver is tagged with a build tag which is used to select which drivers to use at compile time. 