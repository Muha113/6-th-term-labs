#include "mainwindow.h"
#include "menuwindow.h"

#include <QApplication>

int main(int argc, char *argv[])
{
    QApplication a(argc, argv);
    MainWindow w;
    menuwindow f;
    QObject::connect(&w, SIGNAL(newType()), &f, SLOT(getType()));
    QObject::connect(&w, SIGNAL(newColor()), &f, SLOT(getColor()));
    QObject::connect(&w, SIGNAL(newCheckBoxArg()), &f, SLOT(getCheckBoxArg()));
    w.setAttribute(Qt::WA_NoBackground);
    w.show();
    f.show();
    return a.exec();
}
