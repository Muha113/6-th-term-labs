#ifndef MYTHREAD_H
#define MYTHREAD_H

#include <QLineEdit>
#include <QThread>

class mythread : public QThread
{
Q_OBJECT
public:
    explicit mythread(QString, int);
    void run();
signals:
    void send(int n, QString);
private:
    QString name;
    int time;
};

#endif // MYTHREAD_H
