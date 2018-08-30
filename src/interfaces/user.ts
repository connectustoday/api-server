import { prop, Typegoose, ModelType, InstanceType } from 'typegoose';

class IUser extends IAccount {
	@prop({ required: true })
	first_name: string;

	@prop()
	middle_name?: string;

	@prop()
	last_name?: string;

	@prop({ required: true })
	birthday: string;

	@prop()
	gender?: string;

	@prop({ required: true })
	personal_info: IUserProfile;

	@prop()
	experiences?: Array<IExperience>;
}
