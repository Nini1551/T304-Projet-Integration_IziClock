export interface Alarm{
    id: number;
    calendarId: string;
    name: string;
    description: string;
    ringDate: Date;
    preparationTime: number;
    createdAt: Date;
    locationStart: string;
    locationEnd: string;
    ringtone: string;
    transport: string;
    active: boolean;
}

export interface AlarmData{
    ID: number;
    CalendarID: string;
    Name: string;
    Description: string;
    RingDate: string;
    PreparationTime: number;
    CreatedAt: string;
    LocationStart: string;
    LocationEnd: string;
    Ringtone: string;
    Transport: string;
    IsActive: boolean;
}