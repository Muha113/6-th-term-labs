#include "mythread.h"

MyThread::MyThread(QString str)
{
    this->str = str;
}

void MyThread::run()
{
    while(true)
    {
        int rnd = rand() % 100;
        emit send(str + " " + QString::number(rnd));
        sleep(1);
    }
}
