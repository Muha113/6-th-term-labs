#include <iostream>
#include <vector>
#include <map>
#include <algorithm>

using namespace std;

bool hasCycle(vector<pair<int, int>> &cycle, vector<vector<bool>> &basis_matrix, pair<int, int> vertex);
bool hasRowCycle(vector<pair<int, int>> &cycle, vector<vector<bool>> &basis_matrix, pair<int, int> vertex);
bool hasColCycle(vector<pair<int, int>> &cycle, vector<vector<bool>> &basis_matrix, pair<int, int> vertex);

void print_matrix(const vector<vector<int>> &C, int m, int n) {
  for (int i = 0; i < m; i++) {
    for (int j = 0; j < n; j++) {
      cout << C[i][j] << " ";
    }
    cout << "\n";
  }
}

bool hasCycle(vector<pair<int, int>> &cycle, vector<vector<bool>> &basis_matrix, pair<int, int> vertex) {
  return hasRowCycle(cycle, basis_matrix, vertex);
}

bool hasRowCycle(vector<pair<int ,int>> &cycle, vector<vector<bool>> &basis_matrix, pair<int, int> vertex) {
  for(int i = 0; i < basis_matrix[0].size(); i++) {
    if(i == vertex.second || !basis_matrix[vertex.first][i]) continue;
    if(hasColCycle(cycle, basis_matrix, make_pair(vertex.first, i))) {
      cycle.emplace_back(vertex.first, i);
      return true;
    }
  }
  return false;
}

bool hasColCycle(vector<pair<int ,int>> &cycle, vector<vector<bool>> &basis_matrix, pair<int, int> vertex) {
  for(int i = 0; i < basis_matrix.size(); i++) {
    if(i == cycle[0].first && vertex.second == cycle[0].second) {
      cycle.emplace_back(i, vertex.second);
      return true;
    }
    if(i == vertex.first || !basis_matrix[i][vertex.second]) continue;
    if(hasRowCycle(cycle, basis_matrix, make_pair(i, vertex.second))) {
      cycle.emplace_back(i, vertex.second);
      return true;
    }
  }
  return false;
}

void find_potentials(map<int, vector<int>> &i_index_map, map<int, vector<int>> &j_index_map,
                     const vector<vector<int>> &C, vector<vector<bool>> &isVisited, vector<int> &U,
                     vector<int> &V, int i, int j, bool isNeedCalcV) {
  if (!isVisited[i][j]) {
    if (isNeedCalcV) {
      V[j] = C[i][j] - U[i];
    } else {
      U[i] = C[i][j] - V[j];
    }
    isVisited[i][j] = true;
    for (const auto &x: i_index_map[i]) {
      find_potentials(i_index_map, j_index_map, C, isVisited, U, V, i, x, true);
    }
    for (const auto &x: j_index_map[j]) {
      find_potentials(i_index_map, j_index_map, C, isVisited, U, V, x, j, false);
    }
  }
}

void build_basis_plan(const vector<int> &A, const vector<int> &B,
                      vector<vector<int>> &plan, vector<pair<int, int>> &basis) {
  vector<int> A_temp(A);
  vector<int> B_temp(B);
  int prev_size;
  for (int i = 0, j = 0; i < plan.size() && j < plan[0].size();) {
    int mini = min(A_temp[i], B_temp[j]);
    basis.emplace_back(i, j);
    plan[i][j] = mini;
    A_temp[i] -= mini;
    B_temp[j] -= mini;
    if (B_temp[j] == 0) j++;
    if (A_temp[i] == 0) i++;
  }

  prev_size = basis.size();
  for(int i = 0; i < prev_size - 1; i++) {
    if(basis[i].first != basis[i + 1].first && basis[i].second != basis[i + 1].second) {
      basis.emplace_back(basis[i].first, basis[i].second + 1);
    }
  }
  for(int i = 0; i < basis[0].first; i++) {
    basis.emplace_back(i, basis[0].second);
  }
  for(int i = plan.size() - 1; i < basis[prev_size - 1].first; i--) {
    basis.emplace_back(i, basis[prev_size - 1].second);
  }
}

void potential_method(const vector<vector<int>> &C, const vector<int> &A, const vector<int> &B,
                      vector<vector<int>> &plan, vector<pair<int, int>> &basis, int opened_task) {
  vector<int> U(C.size(), 0);
  vector<int> V(C[0].size(), 0);
  vector<vector<bool>> basis_matrix(C.size(), vector<bool>(C[0].size(), false));

  for(const auto &x : basis) {
    basis_matrix[x.first][x.second] = true;
  }

  while(true) {

    vector<vector<bool>> isVisited(C.size(), vector<bool>(C[0].size(), false));

    map<int, vector<int>> i_index_map;
    map<int, vector<int>> j_index_map;

    for (int i = 0; i < basis_matrix.size(); i++) {
      for(int j = 0; j < basis_matrix[0].size(); j++) {
        if(basis_matrix[i][j]) {
          i_index_map[i].emplace_back(j);
          j_index_map[j].emplace_back(i);
        }
      }
    }

    find_potentials(i_index_map, j_index_map, C, isVisited,
                    U, V, U.size() - 1, V.size() - 1, false);

    pair<int, int> min_idx;
    int min_val = INT32_MAX;

    for (int i = 0; i < C.size(); i++) {
      for (int j = 0; j < C[0].size(); j++) {
        if (!basis_matrix[i][j]) {
          int tmp = C[i][j] - U[i] - V[j];
          if (tmp < min_val) {
            min_idx.first = i;
            min_idx.second = j;
            min_val = tmp;
          }
        }
      }
    }

    if (min_val >= 0) {
      print_matrix(plan, opened_task == -1 ? plan.size() - 1 : plan.size(),
              opened_task == 1 ? plan[0].size() - 1 : plan[0].size());
      return;
    }

    vector<pair<int, int>> cycle;
    cycle.emplace_back(min_idx);
    basis_matrix[min_idx.first][min_idx.second] = true;
    if (hasCycle(cycle, basis_matrix, min_idx)) {
      pair<int, int> to_delete;
      min_val = INT32_MAX;

      for (int i = 1; i < cycle.size(); i++) {
        if (i % 2 == 0 && min_val > plan[cycle[i].first][cycle[i].second]) {
          to_delete = cycle[i];
          min_val = plan[cycle[i].first][cycle[i].second];
        }
      }

      for (int i = 1; i < cycle.size(); i++) {
        if (i % 2 == 0) {
          plan[cycle[i].first][cycle[i].second] -= min_val;
        } else {
          plan[cycle[i].first][cycle[i].second] += min_val;
        }
      }

      basis_matrix[to_delete.first][to_delete.second] = false;
    }
  }
}

int main() {
  int m, n;
  cin >> m >> n;
  vector<vector<int>> C(m, vector<int>(n));
  vector<int> A(m);
  vector<int> B(n);
  vector<vector<int>> basis_plan(m, vector<int>(n));
  vector<pair<int, int>> basis;
  int opened_task = 0;

  for (int i = 0; i < m; i++) {
    for (int j = 0; j < n; j++) {
      cin >> C[i][j];
    }
  }
  for (int i = 0; i < m; i++) {
    cin >> A[i];
  }
  for (int i = 0; i < n; i++) {
    cin >> B[i];
  }

  int b_sum = 0, a_sum = 0;
  for(int i : A) {
    a_sum += i;
  }
  for(int i : B) {
    b_sum += i;
  }

  if (a_sum > b_sum) {
    for (int i = 0; i < m; i++) {
      C[i].emplace_back(0);
      basis_plan[i].emplace_back(0);
    }
    B.emplace_back(a_sum - b_sum);
    n++;
    opened_task = 1;
  } else if(a_sum < b_sum) {
    vector<int> add(C[0].size(), 0);
    C.emplace_back(add);
    basis_plan.emplace_back(add);
    A.emplace_back(b_sum - a_sum);
    m++;
    opened_task = -1;
  }
  build_basis_plan(A, B, basis_plan, basis);
  potential_method(C, A, B, basis_plan, basis, opened_task);
}
