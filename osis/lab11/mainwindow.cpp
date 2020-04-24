#include "mainwindow.h"
#include "ui_mainwindow.h"
#include "mythread.h"

MainWindow::MainWindow(QWidget *parent)
    : QMainWindow(parent)
    , ui(new Ui::MainWindow)
{
    ui->setupUi(this);
    mythread *one = new mythread("A", 2);
    mythread *two = new mythread("B", 2);
    mythread *three = new mythread("C", 2);

    connect(one, SIGNAL(send(int, QString)), this, SLOT(update(int, QString)));
    connect(two, SIGNAL(send(int, QString)), this, SLOT(update(int, QString)));
    connect(three, SIGNAL(send(int, QString)), this, SLOT(update(int, QString)));

    one->start();
    two->start();
    three->start();
}

MainWindow::~MainWindow()
{
    delete ui;
}

void MainWindow::update(int i, QString name)
{
    if(name == "A")
        ui->lineEdit1->setText(QString::number(i));
    else if(name == "B")
        ui->lineEdit2->setText(QString::number(i));
    else if(name == "C")
        ui->lineEdit3->setText(QString::number(i));
}

