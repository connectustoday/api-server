import { prop, Typegoose, ModelType, InstanceType } from 'typegoose';
import IAccountSettings from './account-settings';

export default class IUserSettings extends IAccountSettings {

	@prop({ required: true })
	is_full_name_visible: boolean;

	@prop({ required: true })
	blocked_users: Array<string>;

}
