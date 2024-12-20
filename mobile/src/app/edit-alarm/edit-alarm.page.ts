import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { ActivatedRoute } from '@angular/router';
import { Alarm, AlarmData } from 'src/app/interfaces/alarms';
import { AlarmService } from 'src/app/services/alarm.service';
@Component({
  selector: 'app-edit-alarm',
  templateUrl: './edit-alarm.page.html',
  styleUrls: ['./edit-alarm.page.scss'], // Assurez-vous que le fichier SCSS est référencé ici
})
export class EditAlarmePage implements OnInit {
  alarmId: number = 0;
  alarmDetails: AlarmData = {} as AlarmData;
  minDate: string = '';
  initialRingDate: Date = new Date();

  alarmForm!: FormGroup;
  
  constructor(
    private route: ActivatedRoute,
    private alarmService: AlarmService
  ) {}

  doRefresh(event: any) {
    window.location.reload();

    event.target.complete();
  }

  ngOnInit(): void {
    const id = this.route.snapshot.paramMap.get('id');
    if (id) {
      this.alarmId = +id;
      this.getAlarmDetails(this.alarmId);
    }
    this.minDate = this.getCurrentDate();

    this.alarmForm = new FormGroup({
      name: new FormControl(this.alarmDetails.Name, [
        Validators.required, 
        Validators.maxLength(50)
      ]),
      ringDate: new FormControl(this.alarmDetails.RingDate, [
        Validators.required,
      ]),
      preparationTime: new FormControl(this.alarmDetails.PreparationTime, [
        Validators.required,
        Validators.min(0),
        Validators.max(120)
      ]),
      locationStart: new FormControl(this.alarmDetails.LocationStart, [
        Validators.maxLength(100)
      ]),
      locationEnd: new FormControl(this.alarmDetails.LocationEnd, [
        Validators.maxLength(100)
      ]),
      transport: new FormControl('drive'),
      active: new FormControl(this.alarmDetails.IsActive),
    });
  }

  get name() {
    return this.alarmForm.get('name');
  }

  getAlarmDetails(id: number) {
    this.alarmService.getAlarmById(id).subscribe(
      (data: AlarmData) => {
          this.alarmDetails = data;
          this.updateFormValues(data);
          //console.log('Alarm details:', this.alarmDetails);
      },
      (error) => {
        console.error('Error fetching alarm details', error);
      }
    );
  }

  toIsoDateTime(dateTimeString: Date): string {
    const year = dateTimeString.getFullYear();
    const month = (dateTimeString.getMonth() + 1).toString().padStart(2, '0');
    const dateOfMonth = dateTimeString.getDate().toString().padStart(2, '0');
    const hour = dateTimeString.getHours().toString().padStart(2, '0');
    const minute = dateTimeString.getMinutes().toString().padStart(2, '0');
  
    return `${year}-${month}-${dateOfMonth}T${hour}:${minute}:00.000`;
  }

  updateFormValues(data: any) {
    this.alarmForm.patchValue({
      name: data.Name ? data.Name : '',
      ringDate: this.toIsoDateTime(data.RingDate ? new Date(data.RingDate) : new Date()),
      preparationTime: data.PreparationTime ? data.PreparationTime : 0,
      locationStart: data.LocationStart ? data.LocationStart : '',
      locationEnd: data.LocationEnd ? data.LocationEnd : '',
      transport: data.Transport ? data.Transport : 'drive',
      active: data.IsActive ? data.IsActive : false,
    });
  }

  getCurrentDate(): string {
    const today = new Date();
    return today.toISOString().split('T')[0]; 
  }

  onSubmit() {
    const updatedAlarmDetails: AlarmData = {
      Description: '',
      ID: this.alarmDetails.ID,
      CalendarID: this.alarmDetails.CalendarID,
      Name: this.alarmForm.value.name,
      RingDate: new Date(this.alarmForm.value.ringDate).toISOString(),
      PreparationTime: this.alarmForm.value.preparationTime,
      CreatedAt: String(this.alarmDetails.CreatedAt),
      LocationStart: this.alarmForm.value.locationStart,
      LocationEnd: this.alarmForm.value.locationEnd,
      Ringtone: this.alarmDetails.Ringtone,
      Transport: this.alarmForm.value.transport,
      IsActive: this.alarmForm.value.active,
    };

    console.log('Updated alarm details:', updatedAlarmDetails);

    this.alarmService.updateAlarmDetails(updatedAlarmDetails).subscribe(
      (data: AlarmData) => {
        console.log('Response from API:', data); // Affichez la réponse de l'API
        window.location.href = `/alarm-details/${this.alarmId}`; // Redirige et recharge la page
      },
      (error) => {
        console.error('Error updating alarm', error);
      }
    );
  }
}