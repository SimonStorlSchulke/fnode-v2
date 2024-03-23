import { Pipe, PipeTransform } from '@angular/core';

@Pipe({ name: 'filenameFromPath' , standalone: true})

export class FileNameFromPath implements PipeTransform {
  transform(value: string) {
    return value.replace(/^.*[\\/]/, '')
  }
}
