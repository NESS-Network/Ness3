import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { WalletService } from '../../../../services/wallet.service';
import { PurchaseService } from '../../../../services/purchase.service';
import { MatDialogRef } from '@angular/material/dialog';

@Component({
  selector: 'app-add-deposit-address',
  templateUrl: './add-deposit-address.component.html',
  styleUrls: ['./add-deposit-address.component.css'],
})
export class AddDepositAddressComponent implements OnInit {

  form: FormGroup;

  constructor(
    public walletService: WalletService,
    private dialogRef: MatDialogRef<AddDepositAddressComponent>,
    private formBuilder: FormBuilder,
    private purchaseService: PurchaseService,
  ) {}

  ngOnInit() {
    this.initForm();
  }

  generate() {
    this.purchaseService.generate(this.form.value.address, this.form.value.coin_type).subscribe(() => this.dialogRef.close());
  }

  private initForm() {
    this.form = this.formBuilder.group({
      address: ['', Validators.required],
      coin_type: ['', Validators.required],
    });
  }
}
