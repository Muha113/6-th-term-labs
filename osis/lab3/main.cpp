#include <iostream>

using namespace std;

struct node {
    int val;
    node* left;
    node* right;
};

node* new_node(int val) {
    node* n = new node;
    n->val = val;
    n->left = nullptr;
    n->right = nullptr;
    return n;
}

struct tree {
private:
    node* root;
    void updown_print(node* p);
    void clear(node* p);
    node* remove(node* p, int val);
public:
    tree() { root = nullptr; }
    void insert(int val);
    void remove(int val);
    void clear();
    void updown_print();
};

void tree::insert(int val) {
    if(root == nullptr) {
        root = new_node(val);
        return;
    }
    node** cur = &root;
    while((*cur) != nullptr) {
        if((*cur)->val == val) return;
        if(val < (*cur)->val) cur = &(*cur)->left;
        else cur = &(*cur)->right;
    }
    (*cur) = new_node(val);
}

node* tree::remove(node *p, int val) {
    if(p == nullptr)
        return p;

    if(val == p->val) {
        node* tmp;
        if(p->right == nullptr)
            tmp = p->left;
        else {
            node* ptr = p->right;
            if(ptr->left == nullptr){
                ptr->left = p->left;
                tmp = ptr;
            } else {
                node* pmin = ptr->left;
                while(pmin->left != nullptr){
                    ptr  = pmin;
                    pmin = ptr->left;
                }
                ptr->left = pmin->right;
                pmin->left = p->left;
                pmin->right = p->right;
                tmp = pmin;
            }
        }
        delete p;
        return tmp;
    } else if(val < p->val)
        p->left = remove(p->left, val);
    else
        p->right = remove(p->right, val);
    return p;
}

void tree::remove(int val) {
    root = remove(root, val);
}

void tree::clear(node *p) {
    if(p != nullptr) {
        clear(p->left);
        clear(p->right);
        delete(p);
        p = nullptr;
    }
}

void tree::clear() {
    clear(root);
}

void tree::updown_print(node *p) {
    if(p != nullptr) {
        cout << p->val << " ";
        updown_print(p->left);
        updown_print(p->right);
    }
}

void tree::updown_print() {
    cout << "\n";
    updown_print(root);
}

int main() {
    tree t;
    t.insert(3);
    t.insert(5);
    t.insert(1);
    t.insert(4);
    t.updown_print();
    t.remove(1);
    t.updown_print();
    t.clear();
    t.insert(9);
    t.insert(2);
    t.updown_print();
}