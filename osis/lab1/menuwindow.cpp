#include "menuwindow.h"
#include "ui_menuwindow.h"

#include <QMessageBox>

menuwindow::menuwindow(QWidget *parent) :
    QWidget(parent),
    ui(new Ui::menuwindow)
{
    ui->setupUi(this);
    type = 0;
    color = nullptr;
}

menuwindow::~menuwindow()
{
    delete ui;
}

int menuwindow::getType()
{
    return type;
}

QColor menuwindow::getColor()
{
    return color;
}

int menuwindow::getCheckBoxArg()
{
    return ui->drawCheckBox->checkState();
}

void menuwindow::on_rhombusRadioButton_clicked()
{
    type = 1;
}

void menuwindow::on_squareRadioButton_clicked()
{
    type = 2;
}

void menuwindow::on_circleRadioButton_clicked()
{
    type = 3;
}

void menuwindow::on_starRadioButton_clicked()
{
    type = 4;
}

void menuwindow::on_colorRedRadioButton_clicked()
{
    color = Qt::red;
}

void menuwindow::on_colorBlueRadioButton_clicked()
{
    color = Qt::blue;
}

void menuwindow::on_colorGreenRadioButton_clicked()
{
    color = Qt::green;
}
