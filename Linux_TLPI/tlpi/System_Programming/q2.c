//The Ask: If errno is just a global integer, what happens when Thread A fails a system call and Thread B succeeds at the exact same microsecond? Explain the modern implementation of errno that prevents this data race.


//Practical Proof: Write a small C snippet that forces a system call failure. Show the absolutely correct, thread-safe way to capture and print that error to stderr without risking the error code mutating before it is printeid
//

#include<stdio.h>
#include<errno.h>
#include<unistd.h>
#include<fcntl.h>
#include<pthread.h>
#include<string.h>
#include<stdlib.h>
void* thread_creator(void* arg){
	sleep(1);
	int fd = open("app.txt",O_RDONLY|O_CREAT);
	if(fd!=-1){
		printf("thread 1 possibly opened or created app.txt\n");
		close(fd);
	}
	else{
		printf("errno :%d\n",errno);
		printf("error is %s\n",strerror(errno));
	}
	return NULL;
}
void* thread_reader(void* arg){
	int fd = open("app.txt", O_RDONLY);
	if(fd == -1){
		perror("[Thread 2] Expected failure: no app.txt found");
	}
	else{
		printf("thread 2 t2 successfully opened app.txt\n");
		close(fd);
	}
	return NULL;
}
int main(){
	pthread_t t1,t2;

	pthread_create(&t1,NULL,thread_creator,NULL);

	pthread_create(&t2,NULL,thread_reader,NULL);

	pthread_join(t1,NULL);
	pthread_join(t2,NULL);
	return 0;
}
