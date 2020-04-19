#include "mainwindow.h"
#include "ui_mainwindow.h"

MainWindow::MainWindow(QWidget *parent)
    : QMainWindow(parent)
    , ui(new Ui::MainWindow)
{
    ui->setupUi(this);
    ui->drawButton->setStyleSheet("border-image:url(C:/Users/danii/Desktop/SP_LABS/lab8/myimage.jpg)");
    ui->clearButton->setStyleSheet("border-image:url(C:/Users/danii/Desktop/SP_LABS/lab8/myimage2.png)");
}

MainWindow::~MainWindow()
{
    delete ui;
}


void MainWindow::on_drawButton_clicked()
{
    isPenisBig = 1;
    update(0, 0, 904, 623);
    //isPenisBig = false;
}

void MainWindow::on_clearButton_clicked()
{
    isPenisBig = 2;
    update(0, 0, 904, 623);
}

void MainWindow::paintEvent(QPaintEvent *event)
{
    Q_UNUSED(event)
    QPainter painter(this);
    if(isPenisBig == 1)
    {
        painter.setPen(Qt::black);
        painter.drawRect(328, 366, 193, 178); //корпус
        painter.drawRect(420, 392, 85, 61); //окно
        painter.drawRect(351, 419, 49, 125); //дверь
        painter.drawRect(482, 242, 36, 67); //труба
        painter.drawEllipse(729, 237, 80, 80); //солнце
        painter.drawEllipse(28, 451, 179, 52); //пруд
        QPolygon polygon;
        polygon << QPoint(328, 544) << QPoint(267, 511) << QPoint(100, 583)
                << QPoint(611, 610) << QPoint(776, 504) << QPoint(567, 502) << QPoint(521, 544);
        painter.drawPolygon(polygon); //крыльцо
        polygon.clear();
        polygon << QPoint(19, 490) << QPoint(18, 419) << QPoint(23, 480) << QPoint(24, 414) << QPoint(28, 486)
                << QPoint(33, 431) << QPoint(38, 490) << QPoint(39, 419) << QPoint(42, 481) << QPoint(49, 430)
                << QPoint(51, 498) << QPoint(19, 490);
        painter.drawPolygon(polygon); //трава
        polygon.clear();
        polygon << QPoint(270, 366) << QPoint(428, 239) << QPoint(583, 366) << QPoint(270, 366);
        painter.drawPolygon(polygon); //крыша
        polygon.clear();
        painter.drawLine(391, 471, 382, 488); //ручка двери
        painter.drawLine(420, 416, 505, 416);
        painter.drawLine(462, 416, 462, 453);
        painter.drawLine(723, 246, 687, 224);
        painter.drawLine(721, 270, 685, 265);
        painter.drawLine(728, 293, 694, 302);
        painter.drawLine(740, 311, 710, 341);
        painter.drawLine(770, 323, 763, 360);
        painter.drawLine(798, 320, 816, 351);
    }
    else if(isPenisBig == 2)
    {
        painter.eraseRect(0, 200, 903, 423);
    }
    //isPenisBig = 0;
}
