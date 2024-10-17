import { Component, OnInit } from '@angular/core';
import { Alarm } from '../../interfaces/alarms';
import { AlarmService } from 'src/app/services/alarm.service';
import { alarm } from 'src/app/classes/alarms';

@Component({
  selector: 'app-alarms-list',
  templateUrl: './alarms-list.component.html',
  styleUrls: ['./alarms-list.component.scss']
})
export class AlarmsListComponent implements OnInit {
  alarms: Alarm[] = [];

  constructor(private alarmService: AlarmService) {

  }
  
  setAlarms(){
    this.alarmService.getAlarms().subscribe((data: any) => {
      console.log(data);
      for (let alarmData of data) {
        let newAlarm: Alarm = new alarm(alarmData);
        this.alarms.push(newAlarm);
      }
      console.log(this.alarms);
    });
  }

  toggleValue: boolean = false;

  onToggleChanged(event: any) {
    this.toggleValue = event.detail.checked;  // Met à jour la valeur du toggle

    // Ajoutez ici la logique à exécuter lorsque le toggle change
    if (this.toggleValue) {
      console.log(event);
    } else {
      console.log('Le toggle est désactivé');
    }
  }

  ngOnInit() {
    this.setAlarms();
  }

}