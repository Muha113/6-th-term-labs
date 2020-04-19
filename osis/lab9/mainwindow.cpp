#include "mainwindow.h"
#include "ui_mainwindow.h"

MainWindow::MainWindow(QWidget *parent)
    : QMainWindow(parent)
    , ui(new Ui::MainWindow)
{
    ui->setupUi(this);
    que.push("C:/Users/danii/Desktop/SP_LABS/lab9/sprite1.png");
    que.push("C:/Users/danii/Desktop/SP_LABS/lab9/sprite2.png");
    que.push("C:/Users/danii/Desktop/SP_LABS/lab9/sprite3.png");
    que.push("C:/Users/danii/Desktop/SP_LABS/lab9/sprite4.png");
    que.push("C:/Users/danii/Desktop/SP_LABS/lab9/sprite5.png");
    que.push("C:/Users/danii/Desktop/SP_LABS/lab9/sprite6.png");
    offset_x = 0;
    offset_y = 0;
    used_x = used_y = 0;
    current_x = 100;
    current_y = 100;
    isPenisBig = false;
    statement = false;
    timer = new QTimer();
    connect(timer, SIGNAL(timeout()), this, SLOT(Draw()));
    timer->start(150);
}

MainWindow::~MainWindow()
{
    delete ui;
    delete timer;
}

void MainWindow::mousePressEvent(QMouseEvent *event)
{
    Q_UNUSED(event)
    statement = true;
    current_x += used_x;
    current_y += used_y;
    used_x = used_y = 0;
    end_x = event->pos().x();
    end_y = event->pos().y();
    double length = sqrt(pow((end_x - current_x)*1., 2.) + pow((end_y - current_y)*1., 2.));
    offset_x = 5 * (end_x - current_x) / length;
    offset_y = 5 * (end_y - current_y) / length;
}

void MainWindow::paintEvent(QPaintEvent *event)
{
    Q_UNUSED(event)
    QPainter painter(this);
    if(isPenisBig == true)
    {
        QImage img;
        img.load(que.front());
        painter.drawImage(current_x + used_x, current_y + used_y, img);
    }
    else
    {
        painter.eraseRect(0, 0, 800, 600);
    }
}

void MainWindow::Draw()
{
    if(abs(used_x - (end_x - current_x)) < 1 || abs(used_y - (end_y - current_y)) < 1)
    {
        statement = false;
        offset_x = 0;
        offset_y = 0;
    }
    ClearWindow();
    isPenisBig = true;
    update(0, 0, 800, 600);
    que.push(que.front());
    que.pop();
    if(statement)
    {
        used_x += offset_x;
        used_y += offset_y;
    }
}

void MainWindow::ClearWindow()
{
    isPenisBig = false;
    update(0, 0, 800, 600);
}
