#include<stdio.h>
#include<unistd.h>
#include<stdint.h>
#include<time.h>
#define ITER 10000


uint64_t get_ns(){
        struct timespec ts;
        clock_gettime(CLOCK_MONOTONIC,&ts);
        return (uint64_t)ts.tv_sec*1000000000LL + ts.tv_nsec;
}

int main(){
	uint64_t start, end;
	double total_user,per_user;
	double total_sys, per_sys;
        start=get_ns();
	for(int i=0;i<ITER;i++){
		printf("A");
	}
	end=get_ns();
	total_user = (double)(end-start)/1e9;
	per_user = total_user/ITER;
	printf("user call total time: %.6f sec\n", total_user);
        printf("Time per user call: %.12f sec\n",per_user);
	start=get_ns();
	for(int i=0;i<ITER;i++){
		write(STDOUT_FILENO, "A", 1);
	}
	end=get_ns();
	total_sys=(double)(end-start)/1e9;
	per_sys = total_sys/ITER;

	printf("system call total time: %.6f sec\n", total_sys);
        printf("Time per system call: %.12f sec\n",per_sys);
	return 0;
}
      	

