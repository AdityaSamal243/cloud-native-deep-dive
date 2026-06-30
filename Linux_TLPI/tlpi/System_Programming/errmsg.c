#include<stdio.h>
#include "tlpi_hdr.h"
#include<fcntl.h>
#include<errno.h>


int main(){
	int fd = open("heay.txt",O_RDONLY);
	if(fd==-1){
		errMsg("issues found");
	}
	printf("File opened successfully with descriptor %d\n", fd);

    // Always close what you open
       if (close(fd) == -1) {
           errExit("close");
        }

        exit(EXIT_SUCCESS);
}

