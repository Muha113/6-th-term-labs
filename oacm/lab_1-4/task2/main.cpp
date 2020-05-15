#include <iostream>
#include <vector>
#include <algorithm>
#include <map>
#include <iomanip>

using namespace std;

//void print_m_NxN(const vector<vector<double>>& a) {
//    for(int i = 0; i < a.size(); i++) {
//        for (int j = 0; j < a[0].size(); j++)
//            cout << a[i][j] << " ";
//        cout<<"\n";
//    }
//}

void print_vector(const vector<int> &a, int l, int r) {
    for(int i = l; i < r; i++)
        cout<<a[i]<<" ";
    cout<<"\n";
}

void print_vector(const vector<double> &a, int l, int r) {
    for(int i = l; i < r; i++)
        cout<<setprecision(10)<<a[i]<<" ";
    cout<<"\n";
}

vector<double> multiple_matrix_by_vector(vector<vector<double>> &A, vector<double> &B, bool flag) {
    if(flag) { // m by v
        vector<double> res(A.size(), 0);
        for (int i = 0; i < A.size(); i++) {
            for (int j = 0; j < A[0].size(); j++) {
                res[i] += A[i][j] * B[j];
            }
        }
        return res;
    } else { // v by m
        vector<double> res(A[0].size(), 0);
        for(int i = 0; i < A[0].size(); i++) {
            for(int j = 0; j < A.size(); j++) {
                res[i] += A[j][i] * B[j];
            }
        }
        return res;
    }
}

vector<vector<double>> sherman(vector<vector<double>> B, vector<double> Col, int pos) {
    int n = Col.size();
    vector<double> L = multiple_matrix_by_vector(B, Col, true);
    double tmp = L[pos];
    L[pos] = -1.0;
    for(int i = 0; i < L.size(); i++)
        L[i] *= (-1.0 / tmp);
    vector<vector<double>> B_new(n, vector<double>(n, 0));
    for(int i = 0; i < n; i++)
        for(int j = 0; j < n; j++) {
            B_new[j][i] = L[j] * B[pos][i];
            if(j != pos) {
                B_new[j][i] += B[j][i];
            }
        }
    return B_new;
}

vector<vector<double>> inverse(vector<vector<double>> input_matrix) {
    int n = input_matrix.size();
    vector<vector<double>> inverse_matrix(n, vector<double>(n));
    for(int i = 0; i < n; i++)
        inverse_matrix[i][i] = 1.;
    for(int i = 0, j = 0; i < n && j < n; i++, j++) {
        int k = j;
        for(int t = j; t < n; t++) {
            if(abs(input_matrix[t][i]) > abs(input_matrix[k][i]))
                k = t;
        }
        for(int p = 0; p < n; p++)
            swap(inverse_matrix[j][p], inverse_matrix[k][p]);
        for(int p = 0; p < n; p++)
            swap(input_matrix[j][p], input_matrix[k][p]);
        for(int p = 0; p < n; p++) {
            if(p != j) {
                double r = input_matrix[p][i] / input_matrix[j][i];
                for(int g = 0; g < n; g++) {
                    input_matrix[p][g] -= input_matrix[j][g] * r;
                    inverse_matrix[p][g] -= inverse_matrix[j][g] * r;
                }
            }
        }
    }
    for(int i = 0; i < n; i++)
        for(int j = 0; j < n; j++)
            inverse_matrix[i][j] /= input_matrix[i][i];
    return inverse_matrix;
}

int simplex_method_main_phase(int m, int n, vector<vector<double>> A, vector<double> b, vector<double> c, vector<double> &x, vector<int> &basis) {
    vector<vector<double>> A_basis(m, vector<double>(m));
    vector<double> c_basis(m);
    vector<double> j0_col(A.size());
    vector<vector<double>> inversed;
    bool isFirst = true;
    int j0, s0;
    for(int i = 0; i < m; i++) {
        for(int j = 0; j < m; j++) {
            A_basis[j][i] = A[j][basis[i]-1];
        }
    }
    while(true) {
        if(isFirst) {
            inversed = inverse(A_basis);
            isFirst = false;
        } else {
            inversed = sherman(inversed, j0_col, s0);
        }
        for(int i = 0; i < c_basis.size(); i++) {
            c_basis[i] = c[basis[i] - 1];
        }
//        cout<<"C basis :: ";
//        print_vector(c_basis, 0, c_basis.size());
        vector<double> delta;
        auto tmp = multiple_matrix_by_vector(inversed, c_basis, false);
        delta = multiple_matrix_by_vector(A, tmp, false);
//        cout<<"Delta :: ";
//        print_vector(delta, 0, delta.size());
        bool is_all_positive = true;
        for(int i = 0; i < delta.size(); i++) {
            if(delta[i] - c[i] < 0.0 && abs(delta[i] - c[i]) > 1e-6) {
                is_all_positive = false;
                j0 = i;
                break;
            }
        }
        if(!is_all_positive) {
//            vector<double> col(A.size());
            for(int i = 0; i < j0_col.size(); i++)
                j0_col[i] = A[i][j0];
//            cout<<"J0 col :: ";
//            print_vector(j0_col, 0, j0_col.size());
            vector<double> z = multiple_matrix_by_vector(inversed, j0_col, true);
//            cout<<"Z :: ";
//            print_vector(z, 0, z.size());
            double s0_val = numeric_limits<double>::max();
            int ctn = 0;
            for(int i = 0; i < m; i++) {
                if(z[i] > 0.0) {
                    double tmp = x[basis[i] - 1] / z[i];
                    if(tmp < s0_val) {
                        s0_val = tmp;
                        s0 = i;
                    }
                    ctn++;
                }
            }
//            cout<<"S0 val :: "<<s0_val<<"\n";
//            cout<<"S0 :: "<<s0<<"\n";
            if(ctn == 0) {
//                cout<<"Unbounded";
//                exit(0);
                return 0;
            }
//            cout<<"Basis before added j0 :: ";
//            print_vector(basis, 0, basis.size());
            basis[s0] = j0 + 1;
//            cout<<"Basis after added j0 :: ";
//            print_vector(basis, 0, basis.size());
            map<int, pair<bool, int>> basis_map;
            for(int i = 0; i < basis.size(); i++) {
                basis_map[basis[i]].first = true;
                basis_map[basis[i]].second = i;
            }
            for(int i = 0; i < x.size(); i++) {
                if(basis_map[i + 1].first) {
                    if(i == j0)
                        x[i] = s0_val;
                    else
                        x[i] -= s0_val * z[basis_map[i + 1].second];
                } else {
                    x[i] = 0.0;
                }
            }
//            cout<<"X after manipulation :: ";
//            print_vector(x, 0, x.size());
        } else {
//            cout<<"Bounded\n";
//            print_vector(x, 0, x.size());
            return 1;
        }
//        cout<<"-------------------------------------------\n";
    }
}

int simplex_method_first_phase(int m, int n, vector<vector<double>> A, vector<double> b, vector<double> c, vector<double> &x, vector<int> &basis) {
    for(int i = 0; i < b.size(); i++) {
        if(b[i] < 0.0) {
            b[i] *= -1.0;
            for(int j = 0; j < A[0].size(); j++)
                A[i][j] *= -1.0;
        }
    }
    vector<double> x_new(n + m, 0.0);
    for(int i = n; i < n+m; i++)
        x_new[i] = b[i-n];
    vector<double> c_new(n+m, 0.0);
    for(int i = n; i < n+m; i++)
        c_new[i] = -1.0;
    vector<vector<double>> A_new(m, vector<double>(n+m, 0.0));
    for(int i = 0; i < m; i++) {
        for(int j = 0; j < n; j++) {
            A_new[i][j] = A[i][j];
        }
    }
    for(int i = 0; i < m; i++) {
        A_new[i][i+n] = 1.0;
    }
//    vector<int> basis(m);
    for(int i = 0; i < m; i++)
        basis[i] = n+i+1;

    int solution = simplex_method_main_phase(m, m+n, A_new, b, c_new, x_new, basis);
    if(solution == 0) {
//        cout<<"Unbounded";
        return 0;
    } else {
        bool isNoSolution = false;
        for (int i = n; i < n + m; i++) {
            if (x_new[i] != 0.0) {
                isNoSolution = true;
                break;
            }
        }
        if (isNoSolution) {
//            cout << "No solution";
            return -1;
        } else {
//            cout << "Bounded\n";
//            print_vector(x, 0, n);
            for(int i = 0; i < n; i++)
                x[i] = x_new[i];
            return 1;
        }
    }
}

int main() {
    int m, n;
    cin>>m>>n;
    vector<vector<double>> A(m, vector<double>(n));
    vector<double> b(m, 0.0);
    vector<double> c(n, 0.0);
    vector<double> x(n, 0.0);
    vector<int> basis(m, 0);
    for(int i = 0; i < m; i++)
        for(int j = 0; j < n; j++)
            cin>>A[i][j];
    for(int i = 0; i < m; i++)
        cin>>b[i];
    for(int i = 0; i < n; i++)
        cin>>c[i];
//    for(int i = 0; i < n; i++)
//        cin>>x[i];
//    for(int i = 0; i < m; i++)
//        cin>>basis[i];
//    simplex_method_main_phase(m, n, A, b, c, x, basis);
    int solution = simplex_method_first_phase(m, n, A, b, c, x, basis);
    if(solution == -1) {
        cout<<"No solution";
        return 0;
    } else if(solution == 0) {
        cout<<"Unbounded";
        return 0;
    }
//    cout << "\n+++++++++++++++ Main phase ++++++++++++++++++++\n";

    solution = simplex_method_main_phase(m, n, A, b, c, x, basis);
    if(solution == 0) {
        cout<<"Unbounded";
    } else {
        cout<<"Bounded\n";
        print_vector(x, 0, n);
    }
}