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

void MainWindow::mousePressEvent(QMouseEvent *event)
{
    x = event->pos().x();
    y = event->pos().y();
    type = emit this->newType();
    color = emit this->newColor();
    state = emit this->newCheckBoxArg();
    //update(0, 0, this->width(), this->height());
    update();
}

void MainWindow::paintEvent(QPaintEvent *event)
{
    Q_UNUSED(event)
    QPainter painter(this);
    painter.setPen(color);
    if(color != nullptr && type != 0 && state != 0)
    {
        if(type == 1)
        {
            QPolygon polygon;
            polygon << QPoint(x - 50, y) << QPoint(x, y - 100) << QPoint(x + 50, y)
                    << QPoint(x, y + 100) << QPoint(x - 50, y);
            painter.drawPolygon(polygon);
        }
        else if(type == 2)
        {
            painter.drawRect(x - 50, y - 50, 100, 100);
        }
        else if(type == 3)
        {
            painter.drawEllipse(x - 50, y - 50  , 100, 100);
        }
        else if(type == 4)
        {
            QPolygon polygon;
            polygon << QPoint(x - 100, y) << QPoint(x - 30, y - 30) << QPoint(x, y - 100)
                    << QPoint(x + 30, y - 30) << QPoint(x + 100, y) << QPoint(x + 30, y + 30)
                    << QPoint(x, y + 100) << QPoint(x - 30, y + 30) << QPoint(x - 100, y);
            painter.drawPolygon(polygon);
        }
    }
}

