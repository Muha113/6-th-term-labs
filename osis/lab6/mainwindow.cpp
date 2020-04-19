#include "mainwindow.h"
#include "ui_mainwindow.h"

#include <QMessageBox>

MainWindow::MainWindow(QWidget *parent)
    : QMainWindow(parent)
    , ui(new Ui::MainWindow)
{
    ui->setupUi(this);
    timer = new QTimer(this);
    connect(timer, SIGNAL(timeout()), this, SLOT(updateLabel()));
}

MainWindow::~MainWindow()
{
    delete ui;
    delete timer;
}


void MainWindow::on_startButton_clicked()
{
    timer->start(20);
}

void MainWindow::on_stopButton_clicked()
{
    timer->stop();
}

void MainWindow::updateLabel()
{
    if(window()->width() <= ui->myLabel->x() + ui->myLabel->width())
        ui->myLabel->move(ui->myLabel->x() - (window()->width() - ui->myLabel->width()), ui->myLabel->y());
    else
        ui->myLabel->move(ui->myLabel->x() + 3, ui->myLabel->y());
}
