#include<bits/stdc++.h>
using namespace std;

int main(){
    ios_base::sync_with_stdio(false);
    cin.tie(NULL);

    int t;
    cin >> t;
    while(t--){
        int n;
        cin >> n;
        vector<int> arr(n);
        for(int i=0;i<n;i++){
            cin >> arr[i];
        }
        int mini = arr[0];
        int flag = 0;
        for(int i = 1;i<n;i++){
            if(arr[i]<mini){
                flag = 1;
                cout << "NO\n";
                break;
            }
        }
        if(flag==0){
            cout << "yes\n";
        }
    }
    return 0;
}