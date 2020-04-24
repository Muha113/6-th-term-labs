#ifndef MYTHREAD_H
#define MYTHREAD_H

#include <QMutex>
#include <QThread>

class MyThread : public QThread
{
Q_OBJECT
public:
    MyThread(QString str);
    void run();
signals:
    void send(QString);
private:
    QString str;
};

#endif // MYTHREAD_H
