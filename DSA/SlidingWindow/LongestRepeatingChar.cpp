// You are given a string s and an integer k. 
//You can choose any character of the string and change it to any other uppercase English character.
// You can perform this operation at most k times.

//Return the length of the longest substring 
// containing the same letter you can get after performing the above operations.

#include<bits/stdc++.h>
using namespace std;

//BRUTE

int main(){
    //we will try to generate all subarrays
    // in a subarray window we will try to find the frequency of elemnet
    //such that if 
    // window length - maxfreq <=k; find maxlen

    //once  it becomes >k break out. 

    string s;
    getline(cin,s);

    int n = s.length();
    int maxlen=0;
    for(int i=0;i<n;i++){
        // we need to keep track of frequency of elemnts in a window
        // we declare as hash array for this
        int hash[256]={0};
        int maxfreq=0;
        for(int j=i;j<n;j++){
            hash[s[j]]++;
            maxfreq=max(maxfreq,hash[s[j]]);
            if(j-i+1-maxfreq <=k){
                maxlen=max(maxlen,j-i+1);
            }
            else{
                break;
            }
        }
        return maxlen;
    }
}