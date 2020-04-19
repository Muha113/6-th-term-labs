#ifndef MENUWINDOW_H
#define MENUWINDOW_H

#include <QWidget>
#include <QColor>

namespace Ui {
class menuwindow;
}

class menuwindow : public QWidget
{
    Q_OBJECT

public:
    explicit menuwindow(QWidget *parent = nullptr);
    ~menuwindow();

public slots:
    int getType();
    QColor getColor();
    int getCheckBoxArg();

private slots:
    void on_rhombusRadioButton_clicked();

    void on_squareRadioButton_clicked();

    void on_circleRadioButton_clicked();

    void on_starRadioButton_clicked();

    void on_colorRedRadioButton_clicked();

    void on_colorBlueRadioButton_clicked();

    void on_colorGreenRadioButton_clicked();

private:
    Ui::menuwindow *ui;
    QColor color;
    int type;
};

#endif // MENUWINDOW_H
