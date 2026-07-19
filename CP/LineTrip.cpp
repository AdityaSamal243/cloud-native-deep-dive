#include<bits/stdc++.h>
using namespace std;

int main(){
    ios_base::sync_with_stdio(false);
    cin.tie(NULL);
    
    int t;
    cin >> t;
    while(t--){
        int n,x;
        cin >> n >> x;
        vector<int> arr(n);
        for(int i=0;i<n;i++){
            cin >> arr[i];
        }
        int minGas = arr[0];
        for(int i=0;i<n-1;i++){
            minGas = max(minGas,arr[i+1]-arr[i]);
        }
        minGas = max(minGas,2*(x-arr[n-1]));

        cout << minGas << '\n';
    }
    return 0;

}