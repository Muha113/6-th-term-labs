#include "mythread.h"

mythread::mythread(QString name, int time)
{
    this->name = name;
    this->time = time;
}

void mythread::run()
{
    while(true) {
        int r = rand() % 100;
        emit send(r, name);
        sleep(time);
    }
}
