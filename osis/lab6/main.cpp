#include "mainwindow.h"

#include <QApplication>

int main(int argc, char *argv[])
{
    QApplication a(argc, argv);
    a.setOverrideCursor(QCursor(QPixmap("C:/Users/danii/Desktop/SP_LABS/lab6/appcursor.ico")));
    MainWindow w;
    w.show();
    return a.exec();
}
