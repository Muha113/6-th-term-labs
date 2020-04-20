#include <iostream>
#include <vector>
#include <iomanip>

using namespace std;

vector<vector<double>> multiple_matrix(vector<vector<double>> &A, vector<vector<double>> &B) {
    vector<vector<double>> ans(A.size(), vector<double>(B[0].size()));
    for (int i = 0; i < A.size(); ++i) {
        for (int j = 0; j < B[0].size(); ++j) {
            for (int k = 0; k < B.size(); ++k) {
                ans[i][j] += A[i][k] * B[k][j];
            }
        }
    }
    return ans;
}

vector<double> multiple_matrix_vector(vector<vector<double>> &A, vector<double> &B) {
    vector<double> res(B.size(), 0);
    for (int i = 0; i < B.size(); i++) {
        for (int j = 0; j < B.size(); j++) {
            res[i] += A[i][j] * B[j];
        }
    }
    return res;
}

void print_m(vector<vector<double>> A) {
    for(int i = 0; i < A.size(); i++) {
        for (int j = 0; j < A[0].size(); j++)
            cout << setprecision(10) << A[i][j] << " ";
        cout << "\n";
    }
}

int main() {
    int n, k;
    cin>>n>>k;
    k--;
    vector<vector<double>> A(n, vector<double>(n));
    vector<vector<double>> B(n, vector<double>(n));
    vector<double> X(n);
    for(int i = 0; i < n; i++)
        for(int j = 0; j < n; j++)
            cin>>A[i][j];
    for(int i = 0; i < n; i++)
        for(int j = 0; j < n; j++)
            cin>>B[i][j];
    for(int i = 0; i < n; i++)
        cin>>X[i];
    vector<double> L = multiple_matrix_vector(B, X);
    if(L[k] == 0.) {
        cout<<"NO";
        return 0;
    }
    vector<double> L2(n);
    double t = L[k];
    L[k] = -1;
    for(int i = 0; i < n; i++)
        L2[i] = (-1.0 / t) * L[i];
    for(int i = 0; i < n; i++)
        for(int j = 0; j < n; j++) {
            if(i == j) A[i][j] = 1;
            else A[i][j] = 0;
        }
    for(int j = 0; j < n; j++)
        A[j][k] = L2[j];
    vector<vector<double>> B_new = multiple_matrix(A, B);
    cout<<"YES\n";
    print_m(B_new);
}