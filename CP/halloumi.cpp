#include<bits/stdc++.h>
using namespace std;

int main(){

    ios_base::sync_with_stdio(false);
    cin.tie(NULL);
    int t;
    cin >> t ;
    while(t--){
        int n,k;
        cin >> n >>  k ;
        vector<int> arr(n);
        for(int i=0;i<n;i++){
            cin >> arr[i] ;
        }
        if(k>1){
            cout << "YES\n";
        }
        else if(is_sorted(arr.begin(),arr.end())){
            cout << "YES\n";
        }
        else{
            cout << "NO\n";
        }
        
    }
    return 0;
}
