# Overwrite-MBR
An example of how to overwrite the MBR of a hard drive using Golang

Overwriting the systems MBR is a powerful method used by some Malwares to render the sytem unbootalble.
This is because the MBR (first 512 bytes of the hard drive) contains important components to make a system bootable, like the bootloader and partition table.
Since most computers today use UEFI with a GPT partition scheme instead of a legacy BIOS with an MBR partition scheme, this method is no longer as common as it used to be.

- Make sure you run the program with administrator privileges
- Works only with MBR partition scheme (Legacy BIOS)
- Works only on Windows (because I coded it Windows specific)
- Test this program only in a safe environment, like a VM!

## Overwrite-MBR v1
This is a very basic example with no payload. It just overwrites the first 512 bytes of the hard disk with nulls.
