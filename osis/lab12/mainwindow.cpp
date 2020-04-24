#include "mainwindow.h"
#include "ui_mainwindow.h"
#include "mythread.h"

MainWindow::MainWindow(QWidget *parent)
    : QMainWindow(parent)
    , ui(new Ui::MainWindow)
{
    ui->setupUi(this);

    sem.release(1);

    MyThread *thread1 = new MyThread("String1");
    MyThread *thread2 = new MyThread("String2");
    MyThread *thread3 = new MyThread("String3");

    MyThread *thread4 = new MyThread("String4");
    MyThread *thread5 = new MyThread("String5");
    MyThread *thread6 = new MyThread("String6");

    connect(thread1, SIGNAL(send(QString)), this, SLOT(update1(QString)));
    connect(thread2, SIGNAL(send(QString)), this, SLOT(update1(QString)));
    connect(thread3, SIGNAL(send(QString)), this, SLOT(update1(QString)));

    connect(thread4, SIGNAL(send(QString)), this, SLOT(update2(QString)));
    connect(thread5, SIGNAL(send(QString)), this, SLOT(update2(QString)));
    connect(thread6, SIGNAL(send(QString)), this, SLOT(update2(QString)));

    thread1->start();
    thread2->start();
    thread3->start();

    thread4->start();
    thread5->start();
    thread6->start();
}

MainWindow::~MainWindow()
{
    delete ui;
}

void MainWindow::update1(QString str)
{
    mutex.lock();
    ui->lineEdit->setText(str);
    mutex.unlock();
}

void MainWindow::update2(QString str)
{
    sem.acquire(1);
    ui->lineEdit_2->setText(str);
    sem.release(1);
}

