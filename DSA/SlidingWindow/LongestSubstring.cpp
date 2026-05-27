#include<bits/stdc++.h>
using namespace std;

// BRUTE 


// int main(){
//     String s;
//     getline(cin,s);
//     cout << endl;

//     int n = s.length();
//     int maxlen=0;
//     for(int i=0;i<n;i++){
//         //need to keep track whther they occured or not
//         //if they occured update 0 to 1;
//         int hash[256]={0};
//         for(int j=i;j<n;j++){
//             if(hash[s[j]]==1){
//                 break;
//             }
//             maxlen=max(maxlen,j-i+1);
//             hash[s[j]]=1;
//         }
//     }
//     cout << maxlen;
//     return 0;
// }


//BETTER try O(N) solution

// int main(){
//     string s;
//     getline(cin,s);
//     int l=0;
//     int r=0;
//     int maxlne=0;
//     int n = s.length();
//     int hash[256]={0};
//     while(r<n){
//         //checking if it occured before and keep reseting untill repeating is not removed
//         while(hash[s[r]]==1){
//             hash[s[l]]=0;
//             l++;
//         }
//         maxlen=max(maxlen,r-l+1);
//         hash[s[r]]=1;
//         r++;
//     }
//     return maxlen;
// }

//OPTIMAL

int main(){
    string s;
    getline(cin,s);
    int l=0,r=0;
    int n = s.length();
    int maxlen=0;
    // we will store the index here
    int hash[256]={-1};
    while(r<n){
        //checking if the element occured before(would not be equal to -1) and within the range of l and r;
        if(hash[s[r]]!=-1){
            if(hash[s[r]]>=l){
                l=hash[s[r]]+1; //(move l to a new index now the element not in range of l and r)
            }
        }
        hash[s[r]]=r;
        maxlen=max(maxlen,r-l+1);
    }
    return maxlen;

}