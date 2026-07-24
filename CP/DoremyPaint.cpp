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
    
        unordered_map<int,int> um;
        for(int i=0;i<n;i++){
           um[arr[i]]++;
        }
        if(um.size()>2){
            cout << "No\n";
        }
        else if(um.size()==1){
            cout << "Yes\n";
        }
        else{
            auto it = um.begin();
            int c1 = it->second;
            int c2 = next(it)->second;
            if(abs(c1-c2)<=1){
                cout << "Yes\n";
            }
            else{
                cout << "No\n";
            }
        }
    }
    return 0;
}