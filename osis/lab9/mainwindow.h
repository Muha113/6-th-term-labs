#ifndef MAINWINDOW_H
#define MAINWINDOW_H

#include <QMainWindow>
#include <QTimer>
#include <queue>
#include <QPainter>
#include <QMouseEvent>
#include <math.h>

QT_BEGIN_NAMESPACE
namespace Ui { class MainWindow; }
QT_END_NAMESPACE

class MainWindow : public QMainWindow
{
    Q_OBJECT

public:
    MainWindow(QWidget *parent = nullptr);
    ~MainWindow();

protected:
    void mousePressEvent(QMouseEvent *event);
    void paintEvent(QPaintEvent *event);
    void ClearWindow();

private slots:
    void Draw();

private:
    double used_x, used_y;
    Ui::MainWindow *ui;
    QTimer *timer;
    std::queue<QString> que;
    double offset_x;
    double offset_y;
    double current_x;
    double current_y;
    int end_x;
    int end_y;
    bool statement;
    bool isPenisBig;
};
#endif // MAINWINDOW_H
