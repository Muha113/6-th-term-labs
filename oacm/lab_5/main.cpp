#include <iostream>
#include <vector>
#include <map>
#include <algorithm>

using namespace std;

void print_matrix(const vector<vector<int>> &C) {
    for(int i = 0; i < C.size(); i++) {
        for(int j = 0; j < C[0].size(); j++) {
            cout << C[i][j] << "\t";
        }
        cout << "\n";
    }
}

void print_graph(const vector<vector<int>> &g) {
    for(int i = 0; i < g.size(); i++) {
        cout << i << ": ";
        for (int j = 0; j < g[i].size(); j++) {
            cout << g[i][j] << " ";
        }
        cout << "\n";
    }
}

void print_vector(const vector<int> &v) {
    for(int i = 0; i < v.size(); i++) {
        cout << v[i] << " ";
    }
    cout << "\n";
}

void print_basis(const vector<pair<int, int>> &b) {
    for(int i = 0; i < b.size(); i++) {
        cout << "i: " << b[i].first << " j: " << b[i].second << "\n";
    }
}

bool hasCycle(int v, const vector<vector<int>> &graph, vector<char> &cl, vector<int> &p,
        int &cycle_st, int &cycle_end, int from = - 1) {
    cl[v] = 1;
    for(int i = 0; i < graph[v].size(); i++) {
        int to = graph[v][i];
        if(to == from) continue;
        if(cl[to] == 0) {
            p[to] = v;
            if(hasCycle(to, graph, cl, p, cycle_st, cycle_end, v)) return true;
        } else if(cl[to] == 1) {
            cycle_end = v;
            cycle_st = to;
            return true;
        }
    }
    cl[v] = 2;
    return false;
}

void build_graph(vector<vector<int>> &graph, map<int, vector<int>> &i_index_row_map,
        map<int, vector<int>> &j_index_col_map, vector<pair<int, int>> &basis,
        map<pair<int, int>, int> &temp_basis) {
    for(int i = 0; i < basis.size(); i++) {
        temp_basis[basis[i]] = i;
    }
    for(int i = 0; i < basis.size(); i++) {
        for(int j = 0; j < i_index_row_map[basis[i].first].size(); j++) {
            if(i_index_row_map[basis[i].first][j] == basis[i].second) {
                if(j - 1 >= 0) {
                    graph[i].emplace_back(
                            temp_basis[make_pair(basis[i].first, i_index_row_map[basis[i].first][j - 1])]);
                }
                if(j + 1 < i_index_row_map[basis[i].first].size()) {
                    graph[i].emplace_back(
                            temp_basis[make_pair(basis[i].first, i_index_row_map[basis[i].first][j + 1])]);
                }
            }
        }
        for(int j = 0; j < j_index_col_map[basis[i].second].size(); j++) {
            if(j_index_col_map[basis[i].second][j] == basis[i].first) {
                if(j - 1 >= 0) {
                    graph[i].emplace_back(
                            temp_basis[make_pair(j_index_col_map[basis[i].second][j - 1], basis[i].second)]);
                }
                if(j + 1 < j_index_col_map[basis[i].second].size()) {
                    graph[i].emplace_back(
                            temp_basis[make_pair(j_index_col_map[basis[i].second][j + 1], basis[i].second)]);
                }
            }
        }
    }
}

void add_vertex_to_graph(vector<vector<int>> &graph, map<int, vector<int>> &i_index_row_map,
                         map<int, vector<int>> &j_index_col_map, vector<pair<int, int>> &basis,
                         map<pair<int, int>, int> &temp_basis, int i, int j) {
    for(int k = 0; k < i_index_row_map[i].size(); k++) {
//                    cout << "i_index_row_map[i][k] : " << i_index_row_map[i][k] << " j : " << j << "\n";
        if(i_index_row_map[i][k] > j || k == i_index_row_map[i].size() - 1) {
            if(k != i_index_row_map[i].size() - 1) {
                if (k - 1 >= 0) {
                    int tmp = temp_basis[make_pair(i, i_index_row_map[i][k - 1])];
                    graph[basis.size()].emplace_back(tmp);
                    graph[tmp].emplace_back(basis.size());
                }
            }
            if (k < i_index_row_map[i].size()) {
                int tmp = temp_basis[make_pair(i, i_index_row_map[i][k])];
                graph[basis.size()].emplace_back(tmp);
                graph[tmp].emplace_back(basis.size());
            }
            break;
        }
    }
    for(int k = 0; k < j_index_col_map[j].size(); k++) {
        if(j_index_col_map[j][k] > i || k == j_index_col_map[j].size() - 1) {
            if(k != j_index_col_map[j].size() - 1) {
                if (k - 1 >= 0) {
                    int tmp = temp_basis[make_pair(j_index_col_map[j][k - 1], j)];
                    graph[basis.size()].emplace_back(tmp);
                    graph[tmp].emplace_back(basis.size());
                }
            }
            if (k < j_index_col_map[j].size()) {
                int tmp = temp_basis[make_pair(j_index_col_map[j][k], j)];
                graph[basis.size()].emplace_back(tmp);
                graph[tmp].emplace_back(basis.size());
            }
            break;
        }
    }
}

void find_potentials(map<int, vector<int>> &i_index_map, map<int, vector<int>> &j_index_map,
                     const vector<vector<int>> &C, vector<vector<bool>> &isVisited, vector<int> &U,
                     vector<int> &V, int i, int j, bool isNeedCalcV) {
    if(!isVisited[i][j]) {
        if(isNeedCalcV) {
            V[j] = C[i][j] - U[i];
        } else {
            U[i] = C[i][j] - V[j];
        }
        isVisited[i][j] = true;
        for(const auto &x: i_index_map[i]) {
            find_potentials(i_index_map, j_index_map, C, isVisited, U, V, i, x, true);
        }
        for(const auto &x: j_index_map[j]) {
            find_potentials(i_index_map, j_index_map, C, isVisited, U, V, x, j, false);
        }
    }
}

void build_basis_plan(const vector<vector<int>> &C, const vector<int> &A, const vector<int> &B,
        vector<vector<int>> &plan, vector<pair<int, int>> &basis) {
    vector<int> A_temp(A);
    vector<int> B_temp(B);
    map<pair<int, int>, bool> basis_map;
//    map<int, int> i_index_row_ctn;
//    map<int, int> j_index_col_ctn;
    map<int, vector<int>> i_index_row_map;
    map<int, vector<int>> j_index_col_map;
    for(int i = 0, j = 0; i < plan.size() && j < plan[0].size();) {
        int mini = min(A_temp[i], B_temp[j]);
        basis.emplace_back(i, j);
//        i_index_row_ctn[i]++;
//        j_index_col_ctn[j]++;
        i_index_row_map[i].emplace_back(j);
        j_index_col_map[j].emplace_back(i);
        basis_map[basis[basis.size()-1]] = true;
        plan[i][j] = mini;
        A_temp[i] -= mini;
        B_temp[j] -= mini;
        if (B_temp[j] == 0) j++;
        if (A_temp[i] == 0) i++;
    }
//    print_matrix(plan);

    int p = plan.size() + plan[0].size() - 1 - basis.size();
//    cout << "p = " << p << "\n";
    if(p <= 0) return;
    vector<vector<int>> graph(basis.size() + p);
    map<pair<int, int>, int> temp_basis;

    build_graph(graph, i_index_row_map, j_index_col_map, basis, temp_basis);
//    cout << "temp_basis : \n--------------\n";
//    for(const auto &x: temp_basis) {
//        cout << x.second << " : " << x.first.first << " " << x.first.second << "\n";
//    }
//    cout << "i_index_row_map : \n--------------\n";
//    for(const auto &x: i_index_row_map) {
//         cout << x.first << ": ";
//         for(const auto &y: x.second) {
//             cout << y << " ";
//         }
//         cout << "\n";
//    }
//    cout << "\n";
//    cout << "j_index_col_map : \n--------------\n";
//    for(const auto &x: j_index_col_map) {
//        cout << x.first << ": ";
//        for(const auto &y: x.second) {
//            cout << y << " ";
//        }
//        cout << "\n";
//    }
//    cout << "\n";
//    cout << "graph : \n--------------\n";
//    print_graph(graph);
//    cout << "--------------\n";
    for(int i = 0; i < C.size(); i++) {
        if(p == 0) break;
        for(int j = 0; j < C[0].size(); j++) {
            if(p == 0) break;
            if(!basis_map[make_pair(i, j)]) {
                vector<int> anc(basis.size() + 1, -1);
                vector<char> cl(basis.size() + 1, 0);
                sort(i_index_row_map[i].begin(), i_index_row_map[i].end());
                sort(j_index_col_map[j].begin(), j_index_col_map[j].end());
                add_vertex_to_graph(graph, i_index_row_map, j_index_col_map, basis,
                        temp_basis, i, j);
//                cout << "Iteration : " << i << " " << j << "\n";
//                print_graph(graph);
//                cout << "Vertex : " << basis.size() << "\n";
                int cycle_st = -1, cycle_end;
                hasCycle(basis.size(), graph, cl, anc, cycle_st, cycle_end);
                if(cycle_st == -1) {
                    temp_basis[make_pair(i, j)] = basis.size();
                    basis.emplace_back(i, j);
                    i_index_row_map[i].emplace_back(j);
                    j_index_col_map[j].emplace_back(i);
                    p--;
//                    cout << "Iteration if no cycle: " << i << " " << j << "\n";
//                    cout << "i_index_row_map : \n--------------\n";
//                    for(const auto &x: i_index_row_map) {
//                        cout << x.first << ": ";
//                        for(const auto &y: x.second) {
//                            cout << y << " ";
//                        }
//                        cout << "\n";
//                    }
//                    cout << "\n";
//                    cout << "j_index_col_map : \n--------------\n";
//                    for(const auto &x: j_index_col_map) {
//                        cout << x.first << ": ";
//                        for(const auto &y: x.second) {
//                            cout << y << " ";
//                        }
//                        cout << "\n";
//                    }
//                    cout << "\n";
                } else {
                    for(const auto &x: graph[basis.size()]) {
                        graph[x].pop_back();
                    }
                    graph[basis.size()].clear();
                }
            }
        }
    }
}

void potential_method(const vector<vector<int>> &C, const vector<int> &A, const vector<int> &B,
        vector<vector<int>> &plan, vector<pair<int, int>> &basis) {
    int ctn = 0;
    while(ctn < 10) {
        cout << "Basis :\n";
        print_basis(basis);
        cout << "Plan :\n";
        print_matrix(plan);
        cout << "\n";
        cout << ++ctn << "\n";
        vector<vector<int>> expenses(C);
        vector<int> U(C.size(), 0);
        vector<int> V(C[0].size(), 0);
        map<pair<int, int>, bool> basis_map;
        sort(basis.begin(), basis.end());
        V[V.size() - 1] = 0;

        // calc U and V
        map<int, vector<int>> i_index_map;
        map<int, vector<int>> j_index_map;
        vector<vector<bool>> isVisited(C.size(), vector<bool>(C[0].size(), false));
        for(const auto &x: basis) {
            basis_map[x] = true;
            i_index_map[x.first].emplace_back(x.second);
            j_index_map[x.second].emplace_back(x.first);
        }

        vector<vector<int>> graph(basis.size() + 1);
        map<pair<int, int>, int> temp_basis;
        map<int, pair<int, int>> rev_temp_basis;

        build_graph(graph, i_index_map, j_index_map, basis, temp_basis);

        for(const auto &x: temp_basis) {
            rev_temp_basis[x.second] = x.first;
        }

        for(int i = 0; i < basis.size(); i++) {
            temp_basis[basis[i]] = i;
        }

        find_potentials(i_index_map, j_index_map, C, isVisited, U, V, U.size() - 1, V.size() - 1, false);

        pair<int, int> min_idx = {-1, -1};
        int min_val = INT32_MAX;

        for(int i = 0; i < C.size(); i++) {
            for(int j = 0; j < C[0].size(); j++) {
                if(!basis_map[make_pair(i, j)]) {
                    expenses[i][j] = C[i][j] - U[i] - V[j];
                    if(expenses[i][j] < min_val) {
                        min_idx.first = i;
                        min_idx.second = j;
                        min_val = expenses[i][j];
                    }
                } else {
                    expenses[i][j] = C[i][j];
                }
            }
        }
    //    cout << "Expensies : \n";
    //    print_matrix(expenses);
        // answer
        if(min_val >= 0) {
            cout << "SOLUTION FOUND!\n";
            return;
        }

        vector<int> anc(basis.size() + 1, -1);
        vector<char> cl(basis.size() + 1, 0);
//    cout << "graph before insert : \n";
//    print_graph(graph);
        add_vertex_to_graph(graph, i_index_map, j_index_map, basis,
                            temp_basis, min_idx.first, min_idx.second);
        temp_basis[min_idx] = basis.size();
        int cycle_st = -1, cycle_end;
        if (hasCycle(basis.size(), graph, cl, anc, cycle_st, cycle_end)) {
            vector<int> full_cycle;
            full_cycle.emplace_back(cycle_st);
            for (int v = cycle_end; v != cycle_st; v = anc[v]) {
                full_cycle.emplace_back(v);
            }
            reverse(full_cycle.begin(), full_cycle.end());
            vector<pair<int, int>> cycle;
            pair<int, int> tmp1, tmp2, tmp3;
            rev_temp_basis[full_cycle[full_cycle.size() - 1]] = make_pair(min_idx.first, min_idx.second);
            pair<int, int> insert_pos = min_idx;
            min_val = INT32_MAX;
            min_idx = {-1, -1};
            tmp1 = rev_temp_basis[full_cycle[full_cycle.size() - 1]];
            tmp2 = rev_temp_basis[full_cycle[1]];
            tmp3 = rev_temp_basis[full_cycle[0]];
            full_cycle.emplace_back(full_cycle[0]);
            if (tmp1.first != tmp2.first && tmp1.second != tmp2.second) {
                cycle.emplace_back(tmp3);
                if (min_val > plan[tmp3.first][tmp3.second]) {
                    min_val = plan[tmp3.first][tmp3.second];
                    min_idx = tmp3;
                }
            }

//        cout << "Full cycle :\n";
//        for(auto x: full_cycle) {
//            cout << x << " ";
//        }
//        cout << "\n";
            for (int i = 0; i < full_cycle.size() - 2; i++) {
                tmp1 = tmp3;
                tmp3 = tmp2;
                tmp2 = rev_temp_basis[full_cycle[i + 2]];
                if (tmp1.first != tmp2.first && tmp1.second != tmp2.second) {
                    cycle.emplace_back(tmp3);
                    if (min_val > plan[tmp3.first][tmp3.second]) {
                        min_val = plan[tmp3.first][tmp3.second];
                        min_idx = tmp3;
                    }
                }
            }
            cout << "Cycle :\n";
            for(const auto &i: cycle) {
                cout << i.first << " " << i.second << "\n";
            }
            min_val = INT32_MAX;
            int coeff = -1;
            for(const auto &i: cycle) {
                if(min_val >= plan[i.first][i.second] * coeff) {
                    min_val = plan[i.first][i.second] * coeff;
                    min_idx = i;
                }
                coeff *= -1;
            }
            cout << "min val : " << abs(min_val) << "\n";
            bool switcher = false;
            for (const auto &i: cycle) {
                switcher ? plan[i.first][i.second] -= min_val : plan[i.first][i.second] += min_val;
                switcher = !switcher;
            }
//            int del_vertex_num = temp_basis[min_idx];
//            if (graph[del_vertex_num].size() >= 2) {
//                vector<pair<int, int>> positions(graph[del_vertex_num].size());
//                for (int i = 0; i < graph[del_vertex_num].size(); i++) {
//                    positions[i] = (rev_temp_basis[graph[del_vertex_num][i]]);
//                }
//                if (positions[0].first == positions[1].first) {
//                    for (int i = 0; i < graph[graph[del_vertex_num][0]].size(); i++) {
//                        if (graph[graph[del_vertex_num][0]][i] == del_vertex_num) {
//                            graph[graph[del_vertex_num][0]][i] = graph[del_vertex_num][1];
//                        }
//                    }
//                    for (int i = 0; i < graph[graph[del_vertex_num][1]].size(); i++) {
//                        if (graph[graph[del_vertex_num][1]][i] == del_vertex_num) {
//                            graph[graph[del_vertex_num][1]][i] = graph[del_vertex_num][0];
//                        }
//                    }
//                }
//                if (positions[positions.size() - 1].second == positions[positions.size() - 2].second) {
//                    for (int i = 0; i < graph[graph[del_vertex_num][graph.size() - 1]].size(); i++) {
//                        if (graph[graph[del_vertex_num][graph.size() - 1]][i] == del_vertex_num) {
//                            graph[graph[del_vertex_num][graph.size() - 1]][i] = graph[del_vertex_num][graph.size() - 2];
//                        }
//                    }
//                    for (int i = 0; i < graph[graph[del_vertex_num][graph.size() - 2]].size(); i++) {
//                        if (graph[graph[del_vertex_num][graph.size() - 2]][i] == del_vertex_num) {
//                            graph[graph[del_vertex_num][graph.size() - 2]][i] = graph[del_vertex_num][graph.size() - 1];
//                        }
//                    }
//                }
//            }
//            for (int i = 0; i < graph[basis.size()].size(); i++) {
//                for (int j = 0; j < graph[graph[basis.size()][i]].size(); j++) {
//                    if (graph[graph[basis.size()][i]][j] == basis.size()) {
//                        graph[graph[basis.size()][i]][j] = del_vertex_num;
//                    }
//                }
//            }
//            graph[del_vertex_num] = graph[basis.size()];
//            graph[basis.size()].clear();
            for (int i = 0; i < basis.size(); i++) {
                if (basis[i] == min_idx) {
                    basis[i] = insert_pos;
                    break;
                }
            }
//            for (int i = 0; i < i_index_map[min_idx.first].size(); i++) {
//                if (i_index_map[min_idx.first][i] == min_idx.second) {
//                    swap(i_index_map[min_idx.first][i],
//                         i_index_map[min_idx.first][i_index_map[min_idx.first].size() - 1]);
//                    i_index_map[min_idx.first].pop_back();
//                    break;
//                }
//            }
//            for (int i = 0; i < j_index_map[min_idx.second].size(); i++) {
//                if (j_index_map[min_idx.second][i] == min_idx.first) {
//                    swap(j_index_map[min_idx.second][i],
//                         j_index_map[min_idx.second][j_index_map[min_idx.second].size() - 1]);
//                    j_index_map[min_idx.second].pop_back();
//                    break;
//                }
//            }
//            i_index_map[insert_pos.first].emplace_back(insert_pos.second);
//            j_index_map[insert_pos.second].emplace_back(insert_pos.first);
//            sort(i_index_map[min_idx.first].begin(), i_index_map[min_idx.first].end());
//            sort(j_index_map[min_idx.second].begin(), j_index_map[min_idx.second].end());
//            sort(i_index_map[insert_pos.first].begin(), i_index_map[insert_pos.first].end());
//            sort(j_index_map[insert_pos.second].begin(), j_index_map[insert_pos.second].end());
//            temp_basis[insert_pos] = del_vertex_num;
//            temp_basis[min_idx] = -10;
//            rev_temp_basis[del_vertex_num] = insert_pos;
//            rev_temp_basis[basis.size()] = {-1, -1};
//        cout << "Cycle :\n";
//        for(auto x: cycle) {
//            cout << x.first << " " << x.second << "\n";
//        }
//        cout << "graph : \n";
//        print_graph(graph);
//        cout << "Cycle: ";
//        for(auto x: full_cycle) {
//            cout << x << " ";
//        }
//        cout << "\n";
        } else {
            // delete min_idx from val
            cout << "OTSOSITE MOY ZHIRNI PENIS!\n";
            return;
        }
    }

//    cout << "\n";
//    print_matrix(expenses);
//    cout << "\n";
//    print_vector(U);
//    cout << "\n";
//    print_vector(V);
//    cout << "\n";
//    print_basis(basis);
}

int main()
{
    //m = 3, n = 5;
    int m, n;
    cin>>m>>n;
    vector<vector<int>> C(m, vector<int>(n));
    vector<int> A(m);
    vector<int> B(n);
    vector<vector<int>> basis_plan(m, vector<int>(n));
    vector<pair<int, int>> basis;

//    C = {
//            {2, 8, -5, 7, 10},
//            {11, 5, 8, -8, -4},
//            {1, 3, 7, 4, 2}
//    };
//    A = {20, 30, 25};
//    B = {10, 10, 10, 10, 10};
//    basis_plan = {
//            {10, 10, 0, 0, 0, 0},
//            {0, 0, 10, 10, 10, 0},
//            {0, 0, 0, 0, 0, 25},
//    };
//    basis = {
//            {0, 0}, {0, 1}, {0, 2}, {0, 5},
//            {1, 2}, {1, 3}, {1, 4},
//            {2, 5}
//    };

    for(int i = 0; i < m; i++) {
        for(int j = 0; j < n; j++) {
            cin >> C[i][j];
        }
    }
    for(int i = 0; i < m; i++) {
        cin >> A[i];
    }
    for(int i = 0; i < n; i++) {
        cin >> B[i];
    }

    int b_sum = 0, a_sum = 0;
    for(int i = 0; i < A.size(); i++) {
        a_sum += A[i];
    }
    for(int i = 0; i < B.size(); i++) {
        b_sum += B[i];
    }
    cout << "B sum = " << b_sum << " A sum = " << a_sum << "\n";
    // opened (non balanced)
    if(a_sum > b_sum) {
        for(int i = 0; i < m; i++) {
            C[i].push_back(0);
            basis_plan[i].push_back(0);
        }
        B.push_back(a_sum - b_sum);
        n++;
    }
//    print_matrix(basis_plan);
    build_basis_plan(C, A, B, basis_plan, basis);
    cout << "Basis : \n";
    print_basis(basis);
    cout << "\n";
    print_matrix(basis_plan);
    potential_method(C, A, B, basis_plan, basis);
}
