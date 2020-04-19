#ifndef MAINWINDOW_H
#define MAINWINDOW_H

#include <QMainWindow>
#include <QMouseEvent>
#include <QPainter>
#include <QMessageBox>

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
    void drawRectangle(QPainter painter);

private:
    Ui::MainWindow *ui;
    int x;
    int y;
    QColor color;
    int type;
    int state;

private slots:

signals:
    int newType();
    QColor newColor();
    int newCheckBoxArg();
};
#endif // MAINWINDOW_H
