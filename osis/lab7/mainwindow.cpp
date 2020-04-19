#include "mainwindow.h"
#include "ui_mainwindow.h"

MainWindow::MainWindow(QWidget *parent)
    : QMainWindow(parent)
    , ui(new Ui::MainWindow)
{
    ui->setupUi(this);
}

MainWindow::~MainWindow()
{
    delete ui;
}

bool isExistInQListWidget(QListWidget *listWidget, QString text)
{
    QList<QListWidgetItem *> found = listWidget->findItems(text, Qt::MatchStartsWith);
    return found.length() == 0 ? false : true;
}

void MainWindow::on_addButton_clicked()
{
    if(ui->lineEdit->text() != "" && !isExistInQListWidget(ui->listWidget1, ui->lineEdit->text()))
        ui->listWidget1->addItem(ui->lineEdit->text());
}

void MainWindow::on_clearButton_clicked()
{
    ui->listWidget1->clear();
    ui->listWidget2->clear();
}

void MainWindow::on_torightButton_clicked()
{
    QList<QListWidgetItem *> rez =  ui->listWidget1->selectedItems();
    if(rez.length() != 0)
    {
        QString searchText = rez.first()->text();
        if(!isExistInQListWidget(ui->listWidget2, searchText))
            ui->listWidget2->addItem(searchText);
    }
}

void MainWindow::on_deleteButton_clicked()
{
    qDeleteAll(ui->listWidget1->selectedItems());
    qDeleteAll(ui->listWidget2->selectedItems());
}






