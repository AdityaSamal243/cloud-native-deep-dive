//Given a binary array nums and an integer k, 
// return the maximum number of consecutive 1's in the array if you can flip at most k 0's.
//longest subarrays with max zeroes..at most k zeroes

//nums = [1,1,1,0,0,0,1,1,1,1,0], k = 2
// Output: 6

#include<bits/stdc++.h>
using namespace std;

//BRUTE FORCE APPROACH

// int main(){
    //  int n,k ;
    //  cin >>n >> k;
    //  vector<int> v(n);
    //  for(int i=0;i<n;i++){
    //     cin >> v[i];
    //  }
//      int maxlen=0;
//     //generate all subarrays
//      for(int i=0;i<n;i++){
//         int temp=k; // so that for each new iteration a fresh k is stored
//         for(int j=i;j<n;j++){
//             if(v[j]==0){
//                 temp--;
//             }
//             if(temp<0) break;
//             maxlen=max(maxlen,j-i+1);
//         }
//      }
//      return maxlen;
     
// }
// The above solution is taking O(N2) time complexity



// BETTER APPROACH we need to optimize this by O(n).
// we need sliding window for this 
// we will frame this as longest subarray with atmost k zeroes (<=k zeroes)
//as we find 0 increase count 
// expand r till count <=k
//if count > k .....shrink l

int main(){
     int n,k ;
     cin >>n >> k;
     vector<int> v(n);
     for(int i=0;i<n;i++){
        cin >> v[i];
     }

     int l=0;
     int r=0;
     int count=0;
     int maxlen=0;
     while(r<n){
        if(v[r]==0){
            count++;
        }
        while(count>k){
           if(v[l]==0){
            count --;
           }
           l++;
        }
        maxlen=max(maxlen,r-l+1);
        r++;
     }

}