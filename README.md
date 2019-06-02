# transactionProcessor
A command line program built with Golang
This program wathesa particular input directory for write or create changes on a specified file in the directory. once a change event is fired on the file, the content of the file is read and scanned for specific data, once a matching data is found, as many as it is found; the data is completely taken out of the file and written into a new unique file. This process goes on and on , untill the process is exited, either by a fatal error or from the console.

# start program
run program executable on the command line passing the -h flag for usage guide.

This program accepts three flag inputs :
-in : Input directory to watch for file changes(If not given, the current directory is used as the input directory)
-out : Output directory to write each unique file to(If not given, the current directory is used as the output directory)
-b : Backup directory for unique files created (If not given, no backup is saved)
